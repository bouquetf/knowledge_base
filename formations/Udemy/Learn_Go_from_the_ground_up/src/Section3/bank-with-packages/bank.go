package main

import (
	"fmt"

	"example.com/bank-with-packages/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(balanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Can’t continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			fmt.Println("Your deposit amount: ", depositAmount)

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFileName)
		case 3:
			fmt.Println("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can’t withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFileName)
		default:
			fmt.Println("Goodbye!")
			return
		}

		fmt.Println("Thanks for coosing our bank")
	}
}

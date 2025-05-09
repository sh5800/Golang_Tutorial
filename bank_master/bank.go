package main

import (
	"fmt"
	"com.shreyash/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"



func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil{
		fmt.Println("Error")
		fmt.Print(err)
		fmt.Println("-------------")
		// panic("Cannot continue without balance")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("You can reach us on this number",randomdata.PhoneNumber())

	for {

		presentOptions()

		var choice int

		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		// fmt.Println("Your choice is:", choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is:", accountBalance)
		case 2:
			fmt.Print("Your Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("New balance after deposit is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
		case 3:
			fmt.Print("Your Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdraw amount must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Withdraw amount must be less than or equal to account balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("New balance after withdraw is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
		default:
			fmt.Println("Good Bye !")
			fmt.Println("Thank you for using Go Bank!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your current balance is:", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your Deposit Amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Deposit amount must be greater than 0")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("New balance after deposit is: ",accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Your Withdraw Amount: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Withdraw amount must be greater than 0")
		// 		continue
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Withdraw amount must be less than or equal to account balance")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("New balance after withdraw is: ",accountBalance)
		//  } else {
		// 	fmt.Println("Good Bye !")
		// 	break
		// 	}
	}
}

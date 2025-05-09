package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error){
	data,err :=  os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000.0, errors.New("error reading balance file")
	}

	balanceText := string(data)
	balance,err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000.0 , errors.New("error parsing balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()
	if err != nil{
		fmt.Println("Error")
		fmt.Print(err)
		fmt.Println("-------------")
		panic("Cannot continue without balance")
	}
	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("What do you want to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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

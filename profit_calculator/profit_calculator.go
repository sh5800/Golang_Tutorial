package main

import (
	"errors"
	"fmt"
	"os"
)

const financialFile = "financial.txt"

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Error:", err)
		panic("Cannot continue without Revenue")
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Error:", err)
		panic("Cannot continue without Expenses")
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Error:", err)
		panic("Cannot continue without Tax Rate")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeFinancialsToFile(ebt,profit,ratio)


	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid input for " + infoText + " please enter a positive value")
	}
	return userInput, nil
}

func writeFinancialsToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %1f\nProfit: %1f\nRatio: %1f\n", ebt, profit, ratio)
	os.WriteFile(financialFile, []byte(results), 0644)
}
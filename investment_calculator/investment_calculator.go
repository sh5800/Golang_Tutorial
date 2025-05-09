package main

import (
	"fmt"
	"math"
) 

const inflationRate float64 = 2.5

func main() {
	
	var investmentAmount float64
	interestRate:= 5.5
	var years float64 

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Interest Rate: ")
	fmt.Scan(&interestRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)
	
	//futureValue := (investmentAmount) * math.Pow((1 + interestRate/100), (years)) 
	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, interestRate, years)	


	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n",futureRealValue)


	//Different ways to print the output
	//fmt.Println("Future Value: ",futureValue);
	//fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f",futureValue,futureRealValue)
	//fmt.Println("Future Real Value: ",futureRealValue)

	// fmt.Printf(`Future Value: %.2f
	// Future Real Value: %.2f`,futureValue,futureRealValue)

	fmt.Print(formattedFV,formattedFRV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, interestRate, years float64) (fv float64, frv float64) {
	fv = (investmentAmount) * math.Pow((1 + interestRate/100), (years))
	frv = fv / math.Pow(1 + inflationRate/100, years)

	return fv,frv
	// return
}
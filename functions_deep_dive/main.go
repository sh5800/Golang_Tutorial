package main

import "fmt"

func main() {
	res := sumUp(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(res)

	numbers := []int{1,10,15}
	
	anotherSum := sumUp(1,numbers...)

	fmt.Println(anotherSum)
	
}

func sumUp(startingVal int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
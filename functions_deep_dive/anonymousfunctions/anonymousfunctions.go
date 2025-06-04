package anonymousfunctions

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	doubled := createTransformer(2)
	tripled := createTransformer(3)

	doubledNumbers := transformNumbers(&numbers, doubled)
	tripledNumbers := transformNumbers(&numbers,tripled)

	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int{
		return number * factor
	}
}
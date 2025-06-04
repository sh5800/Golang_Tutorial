package functionsasvalues

import "fmt"

type transfromFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5,1,2}



	doubleNum := transformNumbers(&numbers,double)
	tripleNum := transformNumbers(&numbers,triple)

	fmt.Println(doubleNum)
	fmt.Println(tripleNum)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	res1 := transformNumbers(&numbers,transformerFn1)
	res2 := transformNumbers(&moreNumbers,transformerFn2)

	fmt.Println(res1)
	fmt.Println(res2)

}

func transformNumbers(numbers *[]int, transform transfromFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}
	return tNumbers
}

func getTransformerFunction(numbers *[] int) transfromFn {
	if(*numbers)[0] == 1{
		return double
	}else{
		return triple
	}
}

func double (number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
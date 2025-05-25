package lists

import "fmt"


func main(){
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices,5.99)
	fmt.Println(prices)

	prices = prices[1:]
	fmt.Println(prices)

	discountedPrices := []float64{101.99,80.99,25.99}
	prices = append(prices,discountedPrices...)

	fmt.Println(prices)


}

// func main(){
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{24.05,34.05,44.05,54.05}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]

// 	featuredPrices[0] = 64.05

// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(featuredPrices)

// 	fmt.Println(highlightedPrices)

// 	fmt.Println(prices)

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]

// 	fmt.Println(highlightedPrices)

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
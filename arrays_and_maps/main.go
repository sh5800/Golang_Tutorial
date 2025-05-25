package main 

import "fmt"

type floatMap map[string]float64

func(m floatMap) output(){
	fmt.Println(m)
}

func main(){
	userNames := make([]string, 2,5)

	userNames[0] = "Julie"

	userNames = append(userNames,"Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRating := make(floatMap,3)

	courseRating["angular"] = 4.6
	courseRating["react"] = 4.7

	courseRating.output()

	for index,value :=  range userNames{
		fmt.Println(index)
		fmt.Println(value)
	}

	for key,value := range courseRating{
		fmt.Println(key)
		fmt.Println(value)
	}
}
package main

import "fmt"


type Product struct{
	title string
	id int
	price float64
}

func main() {
	hobbies := [3]string{"Coding", "Reading", "Music"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	firstSlice := hobbies[0:2]
	secondSlice := hobbies[:2]

	fmt.Println(firstSlice)
	fmt.Println(secondSlice)

	firstSlice = hobbies[1:3]
	fmt.Println(firstSlice)

	courseGoals := []string{"Learn Go","To master microservices"}
	fmt.Println(courseGoals)

	courseGoals[1] = "To learn a new language"
	courseGoals = append(courseGoals,"To master the GIN framework")

	fmt.Println(courseGoals)

	products := []Product{{title: "Diary 2025",id: 10, price: 450.50},{title: "Linc Pentonic VRT",id: 20, price: 19.50}}
	products = append(products,Product{title: "Sticky Notes",id: 30, price: 55.00})

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
package main

import (
	"fmt"
	"sync"
)


func printSomething(s string, wg *sync.WaitGroup){
	defer wg.Done() //3. Decrement the waitgroup counter by one. Defer ensure that wg.Done() is called and the next line i.e fmt.Println(s) is executed after the goroutine is done.
	fmt.Println(s)
}

func main(){  //the main function is itself a goroutine
	
	var wg sync.WaitGroup  //1. declare a waitgroup

	words := []string{
		"First",
		"Second",
		"Third",
		"Fourth",
		"Fifth",
		"Sixth",
		"Seventh",
		"Eighth",
		"Ninth",
	}
	
	wg.Add(len(words))  //2. add the number of goroutines to the waitgroup or increase the counter to allow execution of a given number of goroutines.

	for i,x := range words{
		go printSomething(fmt.Sprintf("%d: %s",i,x),&wg)
	}

	wg.Wait() //4. Blocks the execution until the value of waitgroup counter is zero.  // in this program waitgroup counter comes to 0 at this very line

	// the go keyword tells the go compiler to run this specific part of the program as a goroutine i.e and independant thread and the go scheduler takes care of it's execution.

	//time.Sleep(1 * time.Second) //not the best way
	
	wg.Add(1) //6. to not have the waitgroup counter become negative

	printSomething("This is the second thing to be printed!",&wg)
	//5. the &wg again decrements the waitgroup counter and it becomes negative i.e -1
}

//the above program executed so quickly that there was no sufficient time for that goroutine to be executed.

//goroutine are very lightweight threads
//they are specific to go
//they can be created using go keyword
//they are very lightweight and consume less memory
//they are managed by the go scheduler
//group of goroutines called coroutines are manged by the go scheduler
package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string){
	defer wg.Done()

	msg = s
}

func main() {
	msg = "Hello, world!"


	wg.Add(2)
	go updateMessage("Hello, Cosmos!")
	go updateMessage("Hello, Universe!")
	wg.Wait()

	fmt.Println(msg)
}

// when we don't add a mutex and execute the above program using go run -race . we get a DATA RACE warning. This is because the goroutines are trying to access the same variable at the same time.
//to fix this we introduce a mutex
//a mutex is a variable that is used to control access to a shared resource
//what a mutex does is that it allows only one goroutine to access a shared resource at a time.
//This is achieved by making using of the Lock() and Unlock() methods of the sync.Mutex type.
//So on lines 20 and 21 we add a mutex variable and then use the Lock() method to lock the mutex before accessing the msg variable and the Unlock() method to unlock the mutex after accessing the msg variable.
//By this way we can make sure that only one goroutine is accessing the msg variable at a time.




// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string,m *sync.Mutex){
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "Hello, world!"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("Hello, Cosmos!",&mutex)
// 	go updateMessage("Hello, Universe!",&mutex)
// 	wg.Wait()

// 	fmt.Println(msg)
// }
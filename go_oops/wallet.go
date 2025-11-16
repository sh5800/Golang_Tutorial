package main

import "fmt"

type Wallet struct{

}

func(w *Wallet) Pay(){
	fmt.Println("Paying with wallet")
}
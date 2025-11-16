package main

import "fmt"

type UPI struct{
	upiId string
}

func(upi *UPI) NewUPI(Id string){
	upi.upiId = Id
}

func(u *UPI) Pay(){
	fmt.Println("Paying with UPI: ",u.upiId)
}
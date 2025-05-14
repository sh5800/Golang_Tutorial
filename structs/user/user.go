package user

import (
	"fmt"
	"time"
	"errors")

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputUserDetails(){ // here using a pointer is not necessary, but it is a good practice to use a pointer receiver when the method mutates the struct
	fmt.Println(u.firstName,u.lastName, u.birthDate)
}

func (u *User) ClearUserName(){  //to mutate values in a struct, we need to use a pointer receiver
	u.firstName = ""
	u.lastName = ""	
}

func New(firstName, lastName, birthDate string) (*User, error){

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First Name, Last Name and Birth Date are mandatory")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	},nil
}
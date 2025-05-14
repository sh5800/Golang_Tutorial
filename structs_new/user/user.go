package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct{
	email string
	password string
	*User 
}

func (u *User) OutputUserDetails(){ // here using a pointer is not necessary, but it is a good practice to use a pointer receiver when the method mutates the struct
	fmt.Println(u.firstName,u.lastName, u.birthDate)
}

func (a Admin) OutputAdminDetails(){
	fmt.Println(a.email,a.password,a.User.firstName,a.User.lastName,a.User.birthDate,a.User.createdAt)
}

func (u *User) ClearUserName(){  //to mutate values in a struct, we need to use a pointer receiver
	u.firstName = ""
	u.lastName = ""	
}

func NewAdmin(email, password string) (Admin){
	return Admin{
		email: email,
		password: password,
		User: &User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
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
package admin

import (
	"errors"
	"fmt"
	"com.shreyash/structs-intro/user"
)

type Admin struct {
	email    string
	password string
	User     *user.User
}

func New(email, password string, User *user.User) (*Admin, error) {
	if email == "" || password == "" || User == nil{
		return nil, errors.New("One or more fields are missing")
	}

	return &Admin{
		email:    email,
		password: password,
		User:     User,
	}, nil
}

func (a *Admin) OutputAdminDetails() { // here using a pointer is not necessary, but it is a good practice to use a pointer receiver when the method mutates the struct
	a.User.OutputUserDetails()
	fmt.Println(a.email, a.password)
	
}

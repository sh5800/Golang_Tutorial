package main

import (
	"fmt"
	"com.shreyash/structs-intro/user" // import the user package
	"com.shreyash/structs-intro/admin"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	userEmail := getUserData("Please enter your email: ")
	userPassword := getUserData("Choose a password: ")

	var appUser *user.User

	appUser,err  := user.New(userFirstName, userLastName, userBirthDate)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	var appAdmin *admin.Admin

	appAdmin,err = admin.New(userEmail,userPassword,appUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	appAdmin.OutputAdminDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
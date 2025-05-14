package main

import (
	"fmt"
	"com.shreyash/structs-new/user" // import the user package
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

	var appAdmin user.Admin

	appAdmin= user.NewAdmin(userEmail,userPassword)
	
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	appAdmin.OutputAdminDetails()

	appAdmin.OutputUserDetails()
	appAdmin.ClearUserName()
	appAdmin.OutputUserDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
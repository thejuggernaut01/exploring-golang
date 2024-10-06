package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (user *User) outputUserDetails () {
	fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)
}

func (user *User) clearUsername () {
	user.firstName = "Aj"
	user.lastName = "Jackson"
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")
	// userCreatedAt := time.Now()

	var appUser = User  {
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthDate,
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
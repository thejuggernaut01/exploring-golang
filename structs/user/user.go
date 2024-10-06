package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (user *User) OutputUserDetails () {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) ClearUsername () {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin  {
	return Admin{
		email,
		password,
		User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("First name, last name, and birth date are required")
	}

	return &User {
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}, nil
}
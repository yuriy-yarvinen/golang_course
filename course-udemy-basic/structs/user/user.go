package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) PrintUserData() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}
func (u *User) ClearUserData() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
}

func New(firstName, lastName, birthdate string) (User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return User{}, errors.New("all fields are required")
	}
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string, user User) (Admin, error) {
	if email == "" || password == "" {
		return Admin{}, errors.New("email and password are required")
	}
	return Admin{
		email:    email,
		password: password,
		User:     user,
	}, nil
}

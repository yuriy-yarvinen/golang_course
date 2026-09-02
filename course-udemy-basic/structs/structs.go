package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	regularUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// ... do something awesome with that gathered data!
	regularUser.PrintUserData()
	regularUser.ClearUserData()
	regularUser.PrintUserData()

	adminUser, err := user.New("Admin", "User", "01/01/1990")
	if err != nil {
		fmt.Println("Error creating admin user:", err)
		return
	}
	admin, err := user.NewAdmin("example@example.com", "password123", adminUser)
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	admin.PrintUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

package main

import (
	"errors"
	"fmt"
)

type MyCustomError struct {
	text     string
	httpCode int
}

func (e MyCustomError) Error() string {
	return e.text
}

func ProcessHTTPPath(path string) error {
	if path != "/valid_path" {
		return MyCustomError{
			text:     "not found",
			httpCode: 404,
		}
	}
	return nil
}

var errUserNotFound = errors.New("user not found")

func main() {
	err := ProcessHTTPPath("/some_path")
	if err != nil {
		var customErr MyCustomError
		if errors.As(err, &customErr) {
			fmt.Printf("Error code: %d, text: %s\n",
				customErr.httpCode, customErr.text)
		} else {
			fmt.Println("unexpected error:", err.Error())
		}
	}

	_, err := getUserData(123)
	if err != nil {
		if errors.Is(err, ErrorNoSuchUser) {
			fmt.Println("No such user!")
			return
		}
		fmt.Println("Error:", err)
	}

	err := errUserNotFound

	fmt.Println(
		err == errUserNotFound,
		errors.Is(err, errUserNotFound),
	)

	wrappedErr := fmt.Errorf("getting user: %w", err)
	fmt.Println("wrappedErr:", wrappedErr)

	fmt.Println(
		wrappedErr == errUserNotFound,
		errors.Is(wrappedErr, errUserNotFound),
	)
}

// ####### user_model.go #######

var ErrorNoSuchUser = errors.New("no such user")

func getUserData(userId int64) (*UserData, error) {
	users, err := selectUsersFromDB(Condition{
		ID: userId,
	})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ErrorNoSuchUser
	}
	return users[0], nil
}

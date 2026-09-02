package main

import (
	"fmt"
	"os"
	"time"
)

func greet(phrase string, channel chan bool) {
	fmt.Println("hello", phrase)
	WriteToFile(phrase)

	channel <- true
}

func WriteToFile(phrase string) {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(phrase + "\n")
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func slowGreent(phrase string, channel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", phrase)
	WriteToFile(phrase)
	channel <- true
	close(channel)

}

func main() {

	testChan := make(chan bool)
	// name = create type channel with value type of value

	go greet("1111", testChan)
	go greet("2222", testChan)
	go slowGreent("3333", testChan)
	go greet("4444", testChan)
	// <-testChan
	// <-testChan
	// <-testChan
	// <-testChan
	for value := range testChan {
		fmt.Println(value)
	}
}

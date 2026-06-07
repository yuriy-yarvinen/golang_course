package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"yarvinen.ru/struct-practice/note"
)

func main() {
	title := getUserInput("Enter note title:")
	content := getUserInput("Enter note content:")

	Note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	Note.PrintNote()
	err = Note.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var input string
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return getUserInput(prompt)
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	if input == "" {
		fmt.Println("Input cannot be empty. Please try again.")
		return getUserInput(prompt)
	}
	return input
}

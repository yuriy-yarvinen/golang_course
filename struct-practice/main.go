package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"yarvinen.ru/struct-practice/noteClass"
	"yarvinen.ru/struct-practice/todoClass"
)

type Saver interface {
	Save() error
}

type Printer interface {
	Print()
}

type Outputter interface {
	Saver
	Printer
}

func main() {
	title := getUserInput("Enter note title:")
	content := getUserInput("Enter note content:")
	text := getUserInput("Enter todo text:")

	Note, err := noteClass.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	Todo, err := todoClass.New(text)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	err = saveAndPrint(Note)
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	err = saveAndPrint(Todo)
	if err != nil {
		fmt.Println("Error saving todo:", err)
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

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}
	return nil
}

func saveAndPrint(data Outputter) error {
	data.Print()
	return saveData(data)
}

func PrintSomeData(data interface{}) {
	switch v := data.(type) {
	case noteClass.Note:
		fmt.Printf("Note Title: %s, Note Content: %s\n", v.Title, v.Content)
	case todoClass.Todo:
		fmt.Printf("Todo Text: %s\n", v.Text)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	fmt.Printf("Data: %v\n", data)
}

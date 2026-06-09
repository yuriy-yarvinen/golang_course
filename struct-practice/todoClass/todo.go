package todoClass

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (n Todo) Print() {
	fmt.Println("Text:", n.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func New(text string) (Todo, error) {

	if text == "" {
		return Todo{}, errors.New("content cannot be empty")
	}

	return Todo{
		Text: text,
	}, nil
}

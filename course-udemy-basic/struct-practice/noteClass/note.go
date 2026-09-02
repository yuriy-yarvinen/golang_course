package noteClass

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Print() {
	fmt.Println("Title:", n.Title)
	fmt.Println("Content:", n.Content)
	fmt.Println("Created At:", n.CreatedAt)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("title cannot be empty")
	}
	if content == "" {

		return Note{}, errors.New("content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

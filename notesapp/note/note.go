package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("invalid input type")
	}

	return &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) ToString(){
	fmt.Printf("Title: %v\n\nContent: %v\n\nCreated At: %v\n", n.Title, n.Content, n.CreatedAt)
}

func (n Note) Save() error{
	fileName := strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")) + ".json"
	jsonData, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
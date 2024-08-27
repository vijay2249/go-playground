package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func NewTodo(content string) (*Todo, error) {

	if content == "" {
		return nil, errors.New("invalid input type")
	}

	return &Todo{
		Text: content,
	}, nil
}

func (todo Todo) ToString(){
	fmt.Printf("Content: %v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
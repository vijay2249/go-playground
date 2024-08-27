package notesapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wedonttrack.org/learn/go/basics/notesapp/note"
)

func getNoteData() (string, string){
	var title, content string
	getUserInput("Note title: ", &title)
	getUserInput("Note Content: ", &content)
	return title, content
}

func Run(){
	title, content := getNoteData()

	n, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	var newNote note.Note = *n
	newNote.ToString()
	err = newNote.Save()
	if err != nil {
		fmt.Println(err)
		panic("saving note failed")
	}
	fmt.Println("note saving is success")
}

func getUserInput(outputText string, varName *string) {
	fmt.Print(outputText)
	// _, err := fmt.Scanln(varName)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic("wrong input format - exiting")
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	*varName = text //assign data to variab;e
}

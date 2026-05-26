package main

import (
	"fmt"
	"github.com/Rafaellinos/note/notestruct"
	"bufio"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

func main() {

	title, content := getNoteData()

	newNote, err := notestruct.New(title, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(newNote.GetTitle())
	saveData(newNote)
	// newNote.Save()

}

// since Note struct has save method, it works
func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

func getNoteData() (string, string) {
	title:= getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}


func getUserInput(prompt string) string {
	// reade long input
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // inform the stop reader

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") //remove last breakline
	text = strings.TrimSuffix(text, "\r") // remove 

	return text
}


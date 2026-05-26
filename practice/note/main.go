package main

import (
	"fmt"
	"github.com/Rafaellinos/note/notestruct"
	"bufio"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()

	newNote, err := notestruct.New(title, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(newNote.GetTitle())
	newNote.Save()

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
	text = strings.TrimSuffix(text, "\r")

	return text
}


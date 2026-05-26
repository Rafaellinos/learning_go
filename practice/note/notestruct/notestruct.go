package notestruct

import (
	"time"
	"errors"
	"os"
	"strings"
	"encoding/json"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}


func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("could not create note")
	}
	return &Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}


func (note *Note) GetTitle() string {
	return note.title
}


func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	jsonEncoded, err := note.MarshalJSON()
	if err != nil {
		return errors.New("error enconding json")
	}
	return os.WriteFile(fileName + ".json", jsonEncoded, 0644)
}

func (note Note) MarshalJSON() ([]byte, error) {
	// struct inline + struct tags (metadata used by json library)
    return json.Marshal(struct {
        Title     string    `json:"title"`
        Content   string    `json:"content"`
        CreatedAt string	`json:"created_at"`
    }{
        Title:     note.title,
        Content:   note.content,
        CreatedAt: note.createdAt.Format(time.DateTime),
    })
}


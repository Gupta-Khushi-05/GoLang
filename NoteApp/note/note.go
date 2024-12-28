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
	Title     string `json:"title"`
	Content   string
	createdAt time.Time
}

func (note Note) Display(){
	fmt.Printf("The Title of note is : %v and the Content is : %v.\n", note.Title, note.Content)
}

func (note Note) Save() error{
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note)
	if(err != nil){
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New (title, content string) (Note, error){
	if(title == "" || content == ""){
		return Note{}, errors.New("Invalid input, please enter the valid input")
	}
	return Note {
		Title: title, 
		Content: content,
		createdAt: time.Now(),
	}, nil
}

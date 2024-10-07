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
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display()  {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

// We don't need a pointer here
// since we're not editing the note
func (note Note) Save() error  {
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)
	
	// the Marshal fn will only extract and convert struct data 
	// that's made publicly available.
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
		
func New(title, content string) (Note, error)  {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}
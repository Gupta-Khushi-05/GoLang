package main

import (
	"bufio"
	"fmt"
	"noteApp/project3/note"
	"os"
	"strings"
)

func main() {
	title, content := OutputValue()
	userNote, err := note.New(title, content)
	if err != nil{
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if(err != nil){
		fmt.Print("Saving failed..")
		return
	}
	fmt.Print("Successfully saved...  ")
}

func OutputValue() (string, string){
	title := UserInput("Enter the title: ")
	content := UserInput("Enter the content: ")

	return title, content
}

func UserInput(text string) (string) {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil{
		return "" 
	}
 value = strings.TrimSuffix(value, "\n")
 value = strings.TrimSuffix(value, "\r")
	
  return value
}


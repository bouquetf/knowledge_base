package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)

	if err != nil {
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo content:")

	return text
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // pour windows

	return text
}

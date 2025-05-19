package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"com.shreyash/notes/note"
	"com.shreyash/notes/todo"
)

type saver interface{
	Save() error
}

// type displayer interface{
// 	Display()
// }

type outputtable interface{
	saver
	Display()
}

// type outputtable interface{
// 	saver
// 	displayer
// }

func main(){

	printSomething("A string value")
	printSomething(1)
	printSomething(2.0)

	title,content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo,err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote,err := note.New(title,content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func printSomething(value interface{}){

     intVal, ok := value.(int)

	 if ok{
		fmt.Println("Integer: ", intVal)
		return
	 }

	 floatVal, ok := value.(float64)

	 if ok{
		fmt.Println("Float: ", floatVal)
		return
	 }

	 stringVal, ok := value.(string)

	 if ok{
		fmt.Println("String: ", stringVal)
		return
	 }

	// switch value.(type){
	// case int:
	// 	fmt.Println("Integer: ",value)
	// case float64:
	// 	fmt.Println("Float64: ",value)
	// case string:
	// 	fmt.Println("String: ",value)		
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)

	if err != nil {
		return err
	}
	return nil
}

func saveData(data saver) error{
	err := data.Save()

	if err != nil {
		fmt.Println("Error in Save Process")
		return err
	}

	fmt.Println("Saved Successfully")
	return nil
}

func getNoteData() (string, string){
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title,content

}

func getUserInput(prompt string) (string){
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')

	if err != nil{
		return ""
	}

	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")

	
	return text
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "wedonttrack.org/learn/go/basics/Basics"
	// "wedonttrack.org/learn/go/basics/Basics/controlStructures"
	// "wedonttrack.org/learn/go/basics/Basics/functions"
	// "wedonttrack.org/learn/go/basics/structs_custom_types"
	// "wedonttrack.org/learn/go/basics/structs_custom_types/custom"
	// "wedonttrack.org/learn/go/basics/notesapp"
	"wedonttrack.org/learn/go/basics/interfaces"
	// "wedonttrack.org/learn/go/basics/interfaces/todo"
	// "wedonttrack.org/learn/go/basics/generics"
	// arraysslicesmaps "wedonttrack.org/learn/go/basics/Arrays_Slices_Maps"
	// functionsDeepDive "wedonttrack.org/learn/go/basics/functions_deepdive"
	"wedonttrack.org/learn/go/basics/concurrency"
)



func main() {
	// fmt.Println("Hi MOM!!")

	// basics.Basics()
	// functions.CalculateRevenue()
	// controlstructures.Controlstructures()
	// structs_custom_types.Start()
	// custom.Start()
	// notesapp.Run()


	// todoText := getUserInput("Enter todo: ")
	// Todo, err := todo.NewTodo(todoText)
	// checkError(err)

	// Todo.ToString()
	// saveData(*Todo)

	// todoText = getUserInput("Enter todo: ")
	// Todo, err = todo.NewTodo(todoText)
	// checkError(err)

	// err = outputAndSaveData(Todo)
	// checkError(err)

	// generics.Main()

	// arraysslicesmaps.Arrays()
	// arraysslicesmaps.Slices()
	// arraysslicesmaps.DeepDiveSlices()
	// arraysslicesmaps.DynamicArrays()
	// arraysslicesmaps.Exercise()
	// arraysslicesmaps.Maps()

	// arraysslicesmaps.UsingMake()
	// functionsDeepDive.Main()
	// functionsDeepDive.AnonymousFunctions()
	// functionsDeepDive.Clousers()
	// functionsDeepDive.Recursions()
	// functionsDeepDive.VariadicFunctions()
	concurrency.Main()

}

func checkError(err error){
	if err != nil {
		fmt.Println(err)
		return
	}
}

func outputAndSaveData(data interfaces.OutputAndSave) error {
	data.ToString()
	return saveData(data)
}

func saveData(data interfaces.Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getUserInput(outputText string) string {
	fmt.Print(outputText)
	// _, err := fmt.Scanln(varName)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		printAnything(err)
		panic("wrong input format - exiting")
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}


//the argument type can be "interface{}"
func printAnything(value any) {

	intVal, ok := value.(int)
	if ok {
		intVal += 1
		fmt.Println("intVal: ", intVal)
	}

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("string: ", value)
	case bool:
		fmt.Println("Boolean: ", value)
	}
	// fmt.Println(value)
}
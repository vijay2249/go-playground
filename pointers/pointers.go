package pointers

import "fmt"

func About(){
	fmt.Println(`
	variables that store value address instead of values
	
	& operator retrieves address of stored value
	
	- can avoid unnecessary duplicate data
	- directly mutate value
	- default value of pointer - nil
	
	
	go has inbuilt garbage collector - running in the background
	`)
}

func Main() {
	var age int
	getUserInput("Enter age: ", &age)
	fmt.Println("Age: ", age)

	// var agePointer *int
	agePointer := &age
	fmt.Println("Age: ", *agePointer)
}


func getUserInput(outputText string, varName *int) {
	fmt.Print(outputText)
	_, err := fmt.Scan(varName)
	if err != nil {
		panic("wrong input format - exiting")
	}
}

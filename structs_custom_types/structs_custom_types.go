package structs_custom_types

import (
	"fmt"

	"wedonttrack.org/learn/go/basics/structs_custom_types/structs"
)


func Start() {
	firstName := getUserData("Firstname: ")
	lastName := getUserData("Lastname: ")
	birthDate := getUserData("Birthdate (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthDate)

	// var user UserData
	testUser := structs.UserData{}
	fmt.Println(testUser)

	user, err := structs.NewUserData(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}


	structs.OutputUserData(*user)
	structs.OutputUserDataViaPointers(user)
	user.ToString()

	user.ModifyUserName()
	user.ToString()


	admin := structs.NewAdmin("test@test.com", "123")
	admin.UserData.ToString()
	admin.UserData.ModifyUserName()
	admin.ToString()


}


func getUserData(promptText string) (value string) {
	fmt.Print(promptText)
	fmt.Scan(&value)
	// fmt.Scanln(&value) //to end input if newline input is provided
	return value
}
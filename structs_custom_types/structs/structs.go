package structs

import (
	"fmt"
	"time"
	"errors"
)

type UserData struct {
	firstName string //need to capitalize the field also to make it accessible
	lastName string
	birthDate string
	createdAt time.Time
}

//struct embedding
type Admin struct {
	email string
	password string
	UserData //this is like inheritence in OOPS
}

func NewAdmin(email, password string) Admin {
	//can add validations as well
	return Admin{
		email: email,
		password: password,
		UserData: UserData{
			firstName: "admin",
			lastName: "user",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

//similar to constructors in other languages
func NewUserData(FirstName, LastName, BirthDate string) (*UserData, error) {
	if FirstName == "" || LastName == "" || BirthDate == "" {
		return nil, errors.New("parameters cannot be empty")
	}
	return &UserData {
		firstName: FirstName,
		lastName: LastName,
		birthDate: BirthDate,
		createdAt: time.Now(),
	}, nil
}

//    u -> receiver argument - this is like argument - this is copy rather than actual object/variable 
func (u UserData) ToString(){
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *UserData) ModifyUserName() {
	(*u).firstName = ""
	(*u).lastName = ""
}

func OutputUserData(u UserData){
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func OutputUserDataViaPointers(u *UserData){
	//both methods are accepted in GO
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt) 
}
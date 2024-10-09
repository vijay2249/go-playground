package models

import (
	"errors"

	"wedonttrack.org/gorest/db"
	"wedonttrack.org/gorest/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (u User) Save() error {
	query := `insert into users(email, password) values (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
  _, err = stmt.Exec(u.Email, hashPassword)
	return err
}

func (u *User) ValidateUser() error {
	query := `select id, password from users where email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return errors.New("no such user found")
	}
	isVaildPassword := utils.CheckHashPassword(retrievePassword, u.Password)
	if !isVaildPassword {
		return errors.New("invalid credentials")
	}
	return nil
}
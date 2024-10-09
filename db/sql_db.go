package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") 
	/*
	if we use 
		DB, err := sql.Open("sqlite3", "api.db") 
		this creates local DB variable which is not reflected in global DB variable
		hence when you run createTables() method this will return null pointer exception as by default pointer value is null and 
		you are trying to run Exec method on null object which as usual will throw error 
	*/
	if err != nil {
		panic("Could not connect to db")
	}

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(3)

	createTables()
}

func createTables() {

	createUsersTable := `
	create table if not exists users(
	  id integer primary key autoincrement,
		email text not null unique,
		password text not null
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		// panic(err.Error())
		panic("could not create users table")
	}


	createEventsTable := `
	 create table if not exists events (
	 	id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT not null,
		description text not null,
		location text not null,
		datetime datetime not null,
		user_id integer,
		foreign key (user_id) references users(id)
	 );
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err.Error())
		panic("could not create events tables in db")
	}

}
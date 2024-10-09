package models

import (
	"time"

	"wedonttrack.org/gorest/db"
)

type Event struct {
	ID int64
	Title string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int64
}


//add it to events variable
func (e *Event) Save() error {
	query := `
	insert into events(title, description, location, datetime, user_id)
	values (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID, 
			&event.Title, 
			&event.Description, 
			&event.Location, 
			&event.DateTime,
			&event.UserID,
		)
		if err != nil { return nil, err}
		events = append(events, event)
	}

	return events, nil
}


func GetEventByID(id int64) (*Event, error) {
	query := "select * from events where id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(
		&event.ID, 
		&event.Title, 
		&event.Description, 
		&event.Location, 
		&event.DateTime,
		&event.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
	update events 
	set title = ?, description = ?, location = ?, datetime = ?
	where id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func DeleteEventByID(id int64) error {
	_, err := GetEventByID(id)
	if err != nil {
		return err
	}
	query := `delete from events where id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
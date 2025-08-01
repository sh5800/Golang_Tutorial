package models

import (
	"time"

	"com.shreyash/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       `binding:"required"`
}

// var events = []Event{}

func (e Event) Save() error {
	query := `
		INSERT INTO events (name,description,location,datetime,user_id)
		VALUES(?,?,?,?,?);
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	e.ID = id
	return err
}

func GetEvents() ([]Event, error) {
	query := "SELECT *FROM events;"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID,&event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error){
	query := "SELECT *FROM events WHERE ID=?;"
	row:= db.DB.QueryRow(query,id)

	var event Event

	err := row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)

	if err != nil{
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, datetime = ?, user_id = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name,event.Description,event.Location,event.DateTime,event.UserID,event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec(event.ID)
	return err
}

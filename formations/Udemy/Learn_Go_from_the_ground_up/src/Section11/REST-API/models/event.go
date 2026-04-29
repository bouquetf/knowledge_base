package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, date_time, user_id) 
		VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * from events WHERE id= ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
		Update events
		SET name = ?, description = ?, location = ?, date_time = ?
		WHERE ID = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (event Event) DeleteEventById() error {
	query := "DELETE FROM events WHERE id = ?"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id=? AND user_id=?"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return nil
}

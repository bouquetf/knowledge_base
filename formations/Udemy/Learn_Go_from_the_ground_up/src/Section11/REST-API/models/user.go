package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	user.ID = userID

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievePassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil
}

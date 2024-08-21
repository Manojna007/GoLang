package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func FetchUserByUsername(username string, db *sqlx.DB) (*User, error) {
	var user User

	err := db.Get(&user, "SELECT id, username, password_hash FROM users WHERE username = ?", username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {

			return nil, nil
		}

		return nil, errors.New("database query error")
	}
	return &user, nil
}

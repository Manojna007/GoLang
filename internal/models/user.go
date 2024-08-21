package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// User represents a user in the system
type User struct {
	ID           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"` // This is the field for storing hashed passwords
}

// FetchUserByUsername retrieves a user by their username
func FetchUserByUsername(username string, db *sqlx.DB) (*User, error) {
	var user User
	// Use a parameterized query to avoid SQL injection
	err := db.Get(&user, "SELECT id, username, password_hash FROM users WHERE username = ?", username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Return a specific error for "user not found"
			return nil, nil
		}
		// For other errors, return a generic error
		return nil, errors.New("database query error")
	}
	return &user, nil
}

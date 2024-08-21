package auth

import (
	"GoAssignment/internal/database"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(userID, password string) (*User, error) {
	var user User
	err := database.DB.Get(&user, "SELECT user_id, password_hash FROM users WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

package auth

type User struct {
	UserID       string `db:"user_id"`
	PasswordHash string `db:"password_hash"`
}

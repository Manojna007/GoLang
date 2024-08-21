package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"GoAssignment/internal/database"
	"GoAssignment/internal/jwtutils"
	"GoAssignment/internal/logger"
	"GoAssignment/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch user by username
	user, err := models.FetchUserByUsername(req.Username, database.DB)
	if err != nil || user == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Directly compare the provided password with the stored password
	if user.PasswordHash != req.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create the JWT claims with the expiration time
	claims := jwtutils.JWTClaims{
		Username: user.Username, // Include the username in the JWT claims
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
		},
	}

	// Generate the token using jwtutils
	token, err := jwtutils.GenerateToken(claims)
	if err != nil {
		logger.Error("Error generating token:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send the token in the response
	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

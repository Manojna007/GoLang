package main

import (
	"fmt"
	"log"
	"net/http"

	"GoAssignment/internal/auth"
	"GoAssignment/internal/config"
	"GoAssignment/internal/database"
	"GoAssignment/internal/health"
	"GoAssignment/internal/logger"
	"GoAssignment/internal/middleware"
	"GoAssignment/internal/students"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize logger
	logger.InitLogger(config.AppConfig.LogFile)

	// Initialize database
	if err := database.InitDB(); err != nil {
		logger.Error("Error initializing database: %v", err)
	}
	defer database.DB.Close()

	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/login", auth.LoginHandler).Methods(http.MethodPost)

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuth)
	protected.HandleFunc("/students", students.GetStudents).Methods(http.MethodGet)
	protected.HandleFunc("/students", students.CreateStudent).Methods(http.MethodPost)
	protected.HandleFunc("/students/{id:[0-9]+}", students.UpdateStudent).Methods(http.MethodPut)
	protected.HandleFunc("/students/{id:[0-9]+}", students.DeleteStudent).Methods(http.MethodDelete)
	protected.HandleFunc("/students/{id:[0-9]+}", students.GetStudentByID).Methods(http.MethodGet)

	// Health check route
	r.HandleFunc("/health", health.HealthCheck).Methods(http.MethodGet)

	// Start server
	serverAddress := fmt.Sprintf(":%d", config.AppConfig.Port)
	logger.Error("Starting server on port %d...", config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}

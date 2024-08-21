package health

import (
	"GoAssignment/internal/database"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthChan := make(chan bool)
	go database.CheckDB(healthChan)

	if !<-healthChan {
		http.Error(w, "DB connection failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"OK"}`))
}

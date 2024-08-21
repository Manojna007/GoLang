package students

import (
	"GoAssignment/internal/database"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var students []Student
	err := database.DB.Select(&students, "SELECT id, name, created_by, created_on, updated_by, updated_on FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	s.CreatedBy = userID
	s.CreatedOn = time.Now()

	_, err := database.DB.NamedExec("INSERT INTO students (name, created_by, created_on) VALUES (:name, :created_by, :created_on)", &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	s.UpdatedBy = userID
	s.UpdatedOn = time.Now()
	s.ID = id

	_, err := database.DB.NamedExec("UPDATE students SET name=:name, updated_by=:updated_by, updated_on=:updated_on WHERE id=:id", &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	_, err := database.DB.Exec("DELETE FROM students WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var s Student
	err := database.DB.Get(&s, "SELECT id, name, created_by, created_on, updated_by, updated_on FROM students WHERE id=?", id)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"myproj/internal/models"
	"myproj/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ProcessTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	remainingTasks, err := services.ProcessPendingTasks(db)
	if err != nil {
		http.Error(w, "Failed to fetch pending tasks", http.StatusInternalServerError)
		return
	}

	if remainingTasks == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All tasks proccessed succesfully"))
	} else {
		http.Error(w, "Some tasks couldn't be processed", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All tasks processed successfully"))
}

func GetTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	tasks, err := models.FetchPendingTasks(db)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"tasks": tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, newTask.Title, newTask.Description, newTask.Completed).Scan(&newTask.ID)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task, err := models.UpdateTask(db, taskID, updatedTask)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update task", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteTask(db, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Task deleted successfully"}`))
}

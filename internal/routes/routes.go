package routes

import (
	"database/sql"
	"myproj/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(db, w, r)
	}).Methods("POST")

	protected := router.PathPrefix("/tasks").Subrouter()
	protected.Use(handlers.JWTMiddleware)

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasks(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTask(db, w, r)
	}).Methods("PUT")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTask(db, w, r)
	}).Methods("DELETE")

	router.HandleFunc("/tasks/process", func(w http.ResponseWriter, r *http.Request) {
		handlers.ProcessTasks(db, w, r)
	}).Methods("POST")

	return router
}

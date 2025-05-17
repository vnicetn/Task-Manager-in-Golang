package main

import (
	loglib "log"
	"myproj/config"
	"myproj/internal/db"
	"myproj/internal/handlers"
	"myproj/internal/routes"
	"myproj/pkg/logger"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	handlers.InitAuth()
	log := logger.NewLogger()
	log.Info("Starting taskmanager application...")

	connStr := config.GetDBConnectionString()
	database, err := db.Init(connStr)
	if err != nil {
		log.Error("Failed to initialize database: %v", err)
		return
	}
	defer database.Close()
	log.Info("Database connected successfully.")

	router := routes.SetupRouter(database)

	log.Info("Server started on :8081")
	loglib.Fatal(http.ListenAndServe(":8081", router))
}

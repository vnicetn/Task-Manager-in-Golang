package db

import (
	"database/sql"
	"myproj/pkg/logger"

	_ "github.com/lib/pq"
)

func Init(connStr string) (*sql.DB, error) {
	log := logger.NewLogger()

	log.Info("Connecting to the database...")
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Failed to connect to database: %v", err)
		return nil, err
	}
	if err = database.Ping(); err != nil {
		log.Error("Database ping failed: %v", err)
		return nil, err
	}

	log.Info("Database connection estabilished.")
	return database, nil
}

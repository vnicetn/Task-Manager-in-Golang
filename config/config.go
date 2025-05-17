package config

import (
	"fmt"
)

func GetDBConnectionString() string {
	user := "postgres"
	password := "123"
	dbname := "taskmanager"
	host := "localhost"
	port := 5432

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, password, dbname, host, port)
}

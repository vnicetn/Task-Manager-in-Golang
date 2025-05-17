package services

import (
	"database/sql"
	"log"
	"myproj/internal/models"
	"time"
)

func ProcessPendingTasks(db *sql.DB) (int, error) {
	tasks, err := models.FetchPendingTasks(db)
	if err != nil {
		return 0, err
	}

	for _, task := range tasks {
		err := models.MarkTaskAsCompleted(db, task.ID)
		if err != nil {
			return 0, err
		}
	}

	return len(tasks), nil
}

func processTask(task models.Task, db *sql.DB) {
	log.Printf("Task with an ID of %d, %s starts processing", task.ID, task.Title)

	time.Sleep(2 * time.Second)

	_, err := db.Exec("UPDATE tasks SET completed = true WHERE id = $1", task.ID)
	if err != nil {
		log.Printf("Task ID %d: Failed to update status to 'Completed': %v", task.ID, err)
		return
	}
	log.Printf("Task ID %d: Status updated to 'Completed'", task.ID)
}

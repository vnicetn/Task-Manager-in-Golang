package models

import "database/sql"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func MarkTaskAsCompleted(db *sql.DB, taskID int) error {
	query := `
		UPDATE tasks
		SET completed = true
		WHERE id = $1
	`

	_, err := db.Exec(query, taskID)
	if err != nil {
		return err
	}

	return nil
}

func FetchPendingTasks(db *sql.DB) ([]Task, error) {
	query := `
		SELECT id, title, description, completed
		FROM tasks
		WHERE completed = false
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func UpdateTask(db *sql.DB, taskID int, updatedTask Task) (Task, error) {
	var task Task

	query := `
	UPDATE tasks
	SET title = $1, description = $2, completed = $3
	WHERE id = $4
	RETURNING id, title, description, completed
	`
	err := db.QueryRow(query, updatedTask.Title, updatedTask.Description, updatedTask.Completed, taskID).
		Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func DeleteTask(db *sql.DB, taskID int) error {
	query := `DELETE FROM tasks WHERE id = 1$`
	result, err := db.Exec(query, taskID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

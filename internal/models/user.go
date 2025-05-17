package models

import "database/sql"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(db *sql.DB, user *User) error {
	query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id
	`
	return db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	var user User

	query := `
		SELECT id, username, password
		FROM users
		WHERE username = $1
	`

	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

package services

import (
	"database/sql"
	"errors"
	"myproj/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err := models.CreateUser(db, user); err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(db *sql.DB, user *models.User, jwtKey string) (string, error) {
	storedUser, err := models.GetUserByUsername(db, user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := generateJWT(storedUser.ID, jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateJWT(userID int, jwtKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat": time.Now().Unix(),                     // Issued at time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtKey))
}

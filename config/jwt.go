package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateJWTKey() string {
	keyLength := 32

	key := make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Failed to generate JWT key: %v", err)
	}

	return base64.StdEncoding.EncodeToString(key)
}

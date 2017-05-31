package utils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"github.com/rymccue/fifa-tracker/models"

	"golang.org/x/crypto/scrypt"
)

// GenerateSalt generates a random salt
func GenerateSalt() string {
	saltBytes := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		log.Fatal(err)
	}
	salt := make([]byte, 32)
	hex.Encode(salt, saltBytes)
	return string(salt)
}

// HashPassword hashes a string
func HashPassword(password, salt string) string {
	hashedPasswordBytes, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		log.Fatal("Unable to hash password")
	}
	hashedPassword := make([]byte, 64)
	hex.Encode(hashedPassword, hashedPasswordBytes)
	return string(hashedPassword)
}

// VerifyPassword verifies if password matches the users password
func VerifyPassword(password string, u *models.User) bool {
	hashedPassword := HashPassword(password, u.Salt)
	return u.Password == hashedPassword
}

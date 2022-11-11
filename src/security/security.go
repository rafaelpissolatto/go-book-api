package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash receive a string and put a into hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare hash and password and return if them are equal
func VerifyPassword(passwordWithHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}

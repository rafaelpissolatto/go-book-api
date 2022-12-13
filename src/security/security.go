package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Hash receive a string and put a into hash
func Hash(password string) ([]byte, error) {
	log.Println("Hashing password...")
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare hash and password and return if them are equal
func VerifyPassword(passwordWithHash, passwordString string) error {
	log.Println("Verifying password...")
	err := bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
	if err != nil {
		return err
	}
	return nil
}

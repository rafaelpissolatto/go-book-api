package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user in the social media
type User struct {
	ID        uint64    `json:"id,omitempty"` //omitempty: if the value is empty, it will not be shown (json)
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare prepares the user to be saved in the database (validate and format)
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

// validate user fields is not empty
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Nickname == "" {
		return errors.New("nickname is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid e-mail format")
	}
	if step == "registration" && user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// format removes the spaces from the extremes of the string
func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
	// For password, we don't need to trim spaces

	if step == "registration" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}
	return nil
}

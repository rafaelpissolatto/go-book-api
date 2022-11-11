package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Represents a repository for users
type users struct {
	db *sql.DB
}

// NewUsersRepository creates a new repository for users
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create creates a user
func (u users) Create(name, nickname, email, password string) (uint64, error) {
	statement, err := u.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(name, nickname, email, password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

// Search searches for users
func (u users) Search(nameOrNick string) ([]models.User, error) {
	// nameOrNick = "%" + nameOrNick + "%"
	// or nameOrNick = "%" + strings.ToLower(nameOrNick) + "%"
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%
	rows, err := u.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE name LIKE ? OR nickname LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u users) SearchByID(userID uint64) (models.User, error) {
	rows, err := u.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// List lists users
func (u users) ListUsers() ([]models.User, error) {
	rows, err := u.db.Query("SELECT id, name, nickname, email, createdAt FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.Password,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Update updates a user in the database
func (u users) Update(userID uint64, user models.User) error {
	statement, err := u.db.Prepare("UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nickname, user.Email, userID); err != nil {
		return err
	}

	return nil
}

// Delete exclude information from the database
func (u users) Delete(userID uint64) error {
	statement, err := u.db.Prepare("DELETE FROM users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// SearchByEmail search a user by email and return his id and his password with hash
func (u users) SearchByEmail(email string) (models.User, error) {
	row, err := u.db.Query("SELECT id, password FROM users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow allows that user follow another user
func (u users) Follow(userID, followerID uint64) error {
	statement, err := u.db.Prepare(
		"INSERT INTO followers (user_id, follower_id) values (?, ?)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// StopFollow allows that user follow another user
func (u users) StopFollow(userID, followerID uint64) error {
	statement, err := u.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

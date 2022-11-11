package database

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Init initializes the database and creates the dabases and tables if they do not exist
func Init() {
	createDatabase()
	createTableUsuarios()
	createTableFollowers()
}

// createDatabase creates the database if it does not exist
func createDatabase() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	sql := `CREATE DATABASE IF NOT EXISTS devbook DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// createTableUsuarios creates the users table if it does not exist
func createTableUsuarios() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
		nickname VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(60) NOT NULL,
		createdAt timestamp default current_timestamp(),
		PRIMARY KEY (id)
	) ENGINE = InnoDB DEFAULT CHARSET = utf8;`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// createTableFollowers creates the followers table if it does not exist
func createTableFollowers() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	sql := `CREATE TABLE IF NOT EXISTS followers(
		user_id INT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		follower_id INT NOT NULL,
		FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
		PRIMARY KEY(user_id, follower_id)
	) ENGINE = InnoDB DEFAULT CHARSET = utf8;`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Connect opens a connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnectionDB)
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Println("[ERROR] Failed to ping the database", err)
		return nil, err
	}

	return db, nil
}

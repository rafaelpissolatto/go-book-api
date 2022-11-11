package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// stringConnectionDB is the string connection to the database
	StringConnectionDB = ""

	// APIPort is the port that the API will run
	ApiPort = 0

	SecretKey = []byte("")
)

// Load loads the configuration from the environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("[ERROR] Failed to load .env file", err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println("[ERROR] Failed to load API_PORT from .env file, using default value 5000")
		ApiPort = 5000
	}

	StringConnectionDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/metrics"
	"api/src/router"
	"api/src/util"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Used to generate the SECRET_KEY (uncomment to regenerate a new one)
func genAappToken() {
	// Generate a new secret key if it is not set
	secretString := os.Getenv("SECRET_KEY")
	if len(secretString) == 0 {
		key := make([]byte, 64)

		if _, err := rand.Read(key); err != nil {
			log.Fatal(err)
		}

		strBase64 := base64.StdEncoding.EncodeToString(key)
		log.Println("[ERROR] The SECRET_KEY is not set, using a random one (please copy and set the token and try to restart the app):", strBase64)
		return
	}

	log.Println("[TRACE] SECRET_KEY:", secretString)
	return
}

func main() {
	// Load logo
	util.Figure()

	// Generate a new secret key if it is not set
	genAappToken()

	// Load configuration
	config.Load()
	log.Println("[TRACE] Database connection string", config.StringConnectionDB)
	log.Println("[INFO] API running on port", config.ApiPort)

	// Setup database
	log.Println("[INFO] Setting up database...")
	database.Init()

	// Load metrics setup
	log.Println("[INFO] Setting up Server metrics...")
	metrics.InitMetrics()

	// Run
	log.Println("[INFO] Running API!")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}

package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"api/src/util"
	"fmt"
	"log"
	"net/http"
)

// Used to generate the SECRET_KEY (uncomment to regenerate a new one)
// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	strBase64 := base64.StdEncoding.EncodeToString(key)
// 	log.Println(strBase64)
// }

func main() {
	util.Figure()

	// Load configuration
	config.Load()
	fmt.Println("[INFO] Database connection string", config.StringConnectionDB)
	fmt.Println("[INFO] API running on port", config.ApiPort)

	// Setup database
	log.Println("[INFO] Setting up database...")
	database.Init()

	log.Println("[INFO] Running API!")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}

package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
	"time"
)

// Logger write information from the request on the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verify if the user doing the request is already authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

// Monitor is a middleware that collects metrics using Prometheus Client
func Monitor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// StartTime is the time when the server started
var StartTime = time.Now()

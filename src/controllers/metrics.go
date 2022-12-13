package controllers

import (
	"api/src/middlewares"
	"api/src/responses"
	"net/http"
	"runtime"
	"time"
)

// Metrics returns the number of goroutines and the uptime of the server
func Metrics(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, map[string]interface{}{
		"goroutines": runtime.NumGoroutine(),
		"uptime":     time.Since(middlewares.StartTime),
	})

}

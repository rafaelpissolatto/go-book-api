package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct that contains the information about a route
type Route struct {
	URI                     string
	Method                  string
	Function                func(w http.ResponseWriter, r *http.Request)
	RequireAuthentification bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.RequireAuthentification {
			r.HandleFunc(route.URI,
				// middlewares.Monitor(middlewares.Logger(middlewares.Authenticate(route.Function)))).Methods(route.Method)
				middlewares.Authenticate(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				// middlewares.Monitor(middlewares.Logger(route.Function))).Methods(route.Method)
				middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	return r
}

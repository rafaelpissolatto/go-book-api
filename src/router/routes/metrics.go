package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeMetrics = Route{
	URI:                     "/metrics",
	Method:                  http.MethodGet,
	Function:                controllers.Metrics,
	RequireAuthentification: false,
}

package routes

import (
	"api/src/controllers"
	"net/http"
)

// Route is a struct that contains the information about a route
var routesUsers = []Route{
	{
		URI:                     "/users",
		Method:                  http.MethodPost,
		Function:                controllers.CreateUser,
		RequireAuthentification: false,
	},
	{
		URI:                     "/users",
		Method:                  http.MethodGet,
		Function:                controllers.SearchUsers,
		RequireAuthentification: true,
	},
	{
		URI:                     "/users/{userId}",
		Method:                  http.MethodGet,
		Function:                controllers.SearchUser,
		RequireAuthentification: true,
	},
	{
		URI:                     "/users/{userId}",
		Method:                  http.MethodPut,
		Function:                controllers.UpdateUser,
		RequireAuthentification: true,
	},
	{
		URI:                     "/users/{userId}",
		Method:                  http.MethodDelete,
		Function:                controllers.DeleteUser,
		RequireAuthentification: true,
	},
	{
		URI:                     "/users/{userId}/follow",
		Method:                  http.MethodPost,
		Function:                controllers.FollowUser,
		RequireAuthentification: true,
	},
	{
		URI:                     "/users/{userId}/stop-follow",
		Method:                  http.MethodPost,
		Function:                controllers.StopFollowUser,
		RequireAuthentification: true,
	},
	{
		URI:                     "/users/{userId}/followers",
		Method:                  http.MethodGet,
		Function:                controllers.SearchFollowers,
		RequireAuthentification: true,
	},
}

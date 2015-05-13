package routes

import (
	"net/http"

	"github.com/daishisystems/payment/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var AllRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handlers.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		handlers.TodoShow,
	},
}

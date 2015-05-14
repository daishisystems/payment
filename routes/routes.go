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
		"GetById",
		"GET",
		"/payment-cards/{id}",
		handlers.GetById,
	},
	Route{
		"GetAll",
		"GET",
		"/payment-cards",
		handlers.GetAll,
	},
	Route{
		"Save",
		"POST",
		"/payment-cards",
		handlers.Save,
	},
}

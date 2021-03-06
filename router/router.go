package router

import (
	"net/http"

	"github.com/couchbaselabs/gocb"
	"github.com/daishisystems/payment/handlers"
	"github.com/daishisystems/payment/logger"
	"github.com/daishisystems/payment/routes"
	"github.com/gorilla/mux"
)

func NewRouter(b *gocb.Bucket) *mux.Router {

	handlers.Init(b)
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes.AllRoutes {
		var handler http.Handler
		handler = route.HandlerFunc

		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

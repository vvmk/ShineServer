package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		// TODO: if valid but time is up
		// TODO: invalidate non-expired JWTs on logout
		if route.Protected {
			handler = jwtMiddleware.Handler(handler)
		}

		handler = Recoverer(Logger(handler, route.Name))

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

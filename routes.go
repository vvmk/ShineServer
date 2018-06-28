package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		false,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
		false,
	},
	Route{
		"Library",
		"GET",
		"/ssrroutine/library/{userId}",
		GetLibrary,
		false,
	},
	Route{
		"User",
		"GET",
		"/ssruser/users/{userId}",
		GetUser,
		false,
	},
	Route{
		"Routine",
		"GET",
		"/ssrroutine/routines/{routineId}",
		GetRoutine,
		false,
	},
}

package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Library",
		"GET",
		"/ssrroutine/library/{userId}",
		GetLibrary,
	},
	Route{
		"User",
		"GET",
		"/ssruser/users/{userId}",
		GetUser,
	},
	Route{
		"Routine",
		"GET",
		"/ssrroutine/routines/{routineId}",
		GetRoutine,
	},
}

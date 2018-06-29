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
		"Login",
		"POST",
		"/login",
		Login,
		false,
	},
	Route{
		"Library",
		"GET",
		"/users/{userId}/library",
		GetLibrary,
		true,
	},
	Route{
		"GetRoutine",
		"GET",
		"/routines/{routineId}",
		GetRoutine,
		false,
	},
	Route{
		"CreateRoutine",
		"POST",
		"/users/{userId}/routines",
		CreateRoutine,
		true,
	},
	Route{
		"ForkRoutine",
		"POST",
		"/users/{userId}/fork/{routineId}",
		ForkRoutine,
		true,
	},
	Route{
		"EditRoutine",
		"PUT",
		"/users/{userId}/routines/{routineId}",
		EditRoutine,
		true,
	},
	Route{
		"DeleteRoutine",
		"DELETE",
		"/users/{userId}/routines/{routineId}",
		DeleteRoutine,
		true,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{userId}",
		GetUser,
		false,
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		CreateUser,
		false,
	},
	Route{
		"EditUser",
		"PUT",
		"/users/{userId}",
		EditUser,
		true,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/users/{userId}",
		DeleteUser,
		true,
	},
}

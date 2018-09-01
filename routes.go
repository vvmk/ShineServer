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
		"Register",
		"POST",
		"/register",
		Register,
		false,
	},
	Route{
		"Confirm",
		"POST",
		"/confirm",
		ConfirmUser,
		false,
	},
	Route{
		"Library",
		"GET",
		"/users/{userId}/library",
		GetLibrary,
		false,
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
	Route{
		"GetProfile",
		"GET",
		"/users/{userId}/profile",
		GetProfile,
		false,
	},
}

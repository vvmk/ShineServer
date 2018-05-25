package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping, %q", html.EscapeString(r.URL.Path))
}

func GetRoutine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	routineId := vars["routineId"]
	fmt.Fprintf(w, "{ Routine %q }", routineId)
}

func GetLibrary(w http.ResponseWriter, r *http.Request) {
	routines := RepoGetAllRoutines()

	library := Library{
		UserId:    1,
		LibraryId: 1,
		Routines:  routines,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(library); err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Fprintf(w, "{ User with id: %q }", userId)
}

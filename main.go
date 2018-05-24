package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

type Routine struct {
	RoutineId     int
	Title         string
	TotalDuration int
	Character     string
	CreatorTag    string
	CreatorId     int
	Creation_date string
	Popularity    int
	Drills        Drills
}
type Routines []Routine

type Drill struct {
	DrillTitle string
	Duration   int
}
type Drills []Drill

type Library struct {
	LibraryId int
	UserId    int
	Routines  Routines
}

type User struct {
	Tag   string
	Email string
	Bio   string
	Main  string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/ssrroutine/routines/{routineId}", HandleGetRoutine)
	router.HandleFunc("/ssrroutine/library/{userId}", HandleGetLibrary)
	router.HandleFunc("/ssruser/users/{userId}", HandleGetUser)

	log.Fatal(http.ListenAndServe(port, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping, %q", html.EscapeString(r.URL.Path))
}

func HandleGetRoutine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	routineId := vars["routineId"]
	fmt.Fprintf(w, "{ Routine %q }", routineId)
}

func HandleGetLibrary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintf(w, "{ Library for user: %q }", userId)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Fprintf(w, "{ User with id: %q }", userId)
}

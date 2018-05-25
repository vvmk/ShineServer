package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const port = ":8080"

type Routine struct {
	RoutineId     int       `json:"routine_id"`
	Title         string    `json:"title"`
	TotalDuration int       `json:"total_duration"`
	Character     string    `json:"character"`
	CreatorTag    string    `json:"creator_tag"`
	CreatorId     int       `json:"creator_id"`
	CreationDate  time.Time `json:"creation_date"`
	Popularity    int       `json:"popularity"`
	Drills        Drills    `json:"drills"`
}
type Routines []Routine

type Drill struct {
	DrillTitle string `json:"drill_title"`
	Duration   int    `json:"duration"`
}
type Drills []Drill

type Library struct {
	LibraryId int      `json:"library_id"`
	UserId    int      `json:"user_id"`
	Routines  Routines `json:"routines"`
}

type User struct {
	Tag   string `json:"tag"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
	Main  string `json:"main"`
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
	routines := Routines{
		Routine{
			RoutineId:     1,
			Title:         "Falco Routine",
			TotalDuration: 30,
			Character:     "Falco",
			CreatorId:     1,
			CreationDate:  time.Now(),
			Popularity:    9,
			Drills: Drills{
				Drill{
					DrillTitle: "Dash->WaveDash",
					Duration:   15,
				},
				Drill{
					DrillTitle: "Short hop Laser",
					Duration:   10,
				},
				Drill{
					DrillTitle: "Dash Dance",
					Duration:   5,
				},
			},
		},
		Routine{
			RoutineId:     2,
			Title:         "Fox Routine",
			TotalDuration: 50,
			Character:     "Fox",
			CreatorId:     1,
			CreationDate:  time.Now(),
			Popularity:    4,
			Drills: Drills{
				Drill{
					DrillTitle: "Dash->WaveDash",
					Duration:   15,
				},
				Drill{
					DrillTitle: "Short hop Laser",
					Duration:   10,
				},
				Drill{
					DrillTitle: "Dash Dance",
					Duration:   5,
				},
				Drill{
					DrillTitle: "Multi-Shine",
					Duration:   20,
				},
			},
		},
		Routine{
			RoutineId:     2,
			Title:         "Justice Routine",
			TotalDuration: 50,
			Character:     "Captain Falcon",
			CreatorId:     1,
			CreationDate:  time.Now(),
			Popularity:    8,
			Drills: Drills{
				Drill{
					DrillTitle: "Dash Dance",
					Duration:   15,
				},
				Drill{
					DrillTitle: "Dthrow->Stomp",
					Duration:   10,
				},
				Drill{
					DrillTitle: "Moonwalk",
					Duration:   25,
				},
			},
		},
	}

	library := Library{
		UserId:    1,
		LibraryId: 1,
		Routines:  routines,
	}
	if err := json.NewEncoder(w).Encode(library); err != nil {
		panic(err)
	}
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Fprintf(w, "{ User with id: %q }", userId)
}

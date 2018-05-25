package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

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

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Fprintf(w, "{ User with id: %q }", userId)
}

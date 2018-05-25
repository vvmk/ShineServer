package main

import (
	"fmt"
	"time"
)

var currentId int

var routines Routines

func init() {
	RepoCreateRoutine(Routine{
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
	})
	RepoCreateRoutine(Routine{
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
	})
	RepoCreateRoutine(Routine{
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
	})
}

func RepoFindRoutine(id int) Routine {
	for _, r := range routines {
		if r.RoutineId == id {
			return r
		}
	}

	return Routine{}
}

func RepoGetAllRoutines() Routines {
	return routines
}

func RepoCreateRoutine(r Routine) Routine {
	currentId += 1
	r.RoutineId = currentId
	routines = append(routines, r)
	return r
}

func RepoDeleteRoutine(id int) error {
	for i, r := range routines {
		if r.RoutineId == id {
			routines = append(routines[:i], routines[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Routine with id of %d to delete", id)
}

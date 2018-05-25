package main

import "time"

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
	UserId int    `json:"user_id"`
	Tag    string `json:"tag"`
	Email  string `json:"email"`
	Bio    string `json:"bio"`
	Main   string `json:"main"`
}

package models

import "time"

type Routine struct {
	RoutineId         int       `json:"routine_id"`
	Title             string    `json:"title"`
	TotalDuration     int       `json:"total_duration"`
	Character         string    `json:"character"`
	OriginalCreatorId int       `json:"original_creator_id"`
	CreatorId         int       `json:"creator_id"`
	Created           time.Time `json:"created"`
	Popularity        int       `json:"popularity"`
	Drills            []Drill   `json:"drills"`
}

type Drill struct {
	DrillTitle string `json:"drill_title"`
	Duration   string `json:"duration"`
}

package models

import (
	"encoding/json"
	"time"
)

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
	Duration   int    `json:"duration"`
}

func (db *DB) FindRoutineById(routineId int) (*Routine, error) {
	var r Routine

	query := `SELECT * FROM routines WHERE routine_id=$1;`

	var d []byte
	err := db.QueryRow(query, routineId).Scan(&r.RoutineId, &r.Title, &r.TotalDuration, &r.Character, &r.CreatorId, &r.Created, &r.Popularity, &d, &r.OriginalCreatorId)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(d, &r.Drills)
	if err != nil {
		panic(err)
	}

	return &r, nil
}

func (db *DB) FindRoutinesByCreator(creatorId int) ([]*Routine, error) {

	query := `SELECT * FROM routines WHERE creator_id = $1;`

	rows, err := db.Query(query, creatorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	routines := make([]*Routine, 0)
	for rows.Next() {
		r := &Routine{}

		var d []byte
		err := rows.Scan(&r.RoutineId, &r.Title, &r.TotalDuration, &r.Character, &r.CreatorId, &r.Created, &r.Popularity, &d, &r.OriginalCreatorId)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(d, &r.Drills)
		if err != nil {
			return nil, err
		}

		routines = append(routines, r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return routines, nil
}

func (db *DB) CreateRoutine(r *Routine) (int, error) {

	query := `INSERT INTO routines(title, total_duration, character, original_creator_id, creator_id, drills)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING routine_id;`

	var routineId int

	drills, err := json.Marshal(r.Drills)
	if err != nil {
		return -1, err
	}

	err = db.QueryRow(query, r.Title, r.TotalDuration, r.Character, r.OriginalCreatorId, r.CreatorId, drills).Scan(&routineId)
	if err != nil {
		return -1, err
	}

	return routineId, nil
}

func (db *DB) UpdateRoutine(r *Routine) (int, error) {
	var routineId int

	query := `UPDATE routines
	SET title = $2, total_duration = $3, character = $4, drills = $5
	RETURNING routine_id;`

	err := db.QueryRow(query, r.Title, r.TotalDuration, r.Character, r.Drills).Scan(&routineId)
	if err != nil {
		return -1, err
	}

	return routineId, nil
}

func (db *DB) DeleteRoutine(routineId int) error {
	query := `DELETE FROM routines WHERE routine_id = $1;`

	_, err := db.Exec(query, routineId)
	if err != nil {
		return nil
	}

	return nil
}

func (db *DB) GetAllRoutines() ([]*Routine, error) {

	query := `SELECT * FROM routines`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	routines := make([]*Routine, 0)
	for rows.Next() {
		r := &Routine{}

		err := rows.Scan(&r.RoutineId, &r.Title, &r.TotalDuration, &r.Character, &r.OriginalCreatorId, &r.CreatorId, &r.Created, &r.Popularity, &r.Drills)
		if err != nil {
			return nil, err
		}

		routines = append(routines, r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return routines, nil
}

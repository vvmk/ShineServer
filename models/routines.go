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

func (db *DB) FindRoutineById(routineId int) (*Routine, error) {
	var routine Routine

	query := `SELECT * FROM routines WHERE routine_id=$1;`

	err := db.QueryRow(query, routineId).Scan(&routine)
	if err != nil {
		return nil, err
	}

	return &routine
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

		// TODO: Drills will definitely break, its a jsonb postgres type
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

func (db *DB) FindRoutinesByLibrary(userId int) ([]*Routine, error) {

}

func (db *DB) CreateRoutine(r *Routine) (int, error) {

}

func (db *DB) AddRoutineToLibrary(userId int, routineId int) error {

}

func (db *DB) RemoveRoutineFromLibrary(userId int, routineId int) error {

}

func (db *DB) UpdateRoutine(r *Routine) (int, error) {

}

func (db *DB) DeleteRoutine(routineId int) error {

}

func (db *DB) GetAllRoutines() ([]*Routine, error) {

}

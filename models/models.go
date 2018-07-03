package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Datastore interface {
	FindRoutineById(routineId int) (*Routine, error)
	FindRoutinesByCreator(creatorId int) ([]*Routine, error)
	CreateRoutine(r *Routine) (int, error)
	UpdateRoutine(r *Routine) (int, error)
	DeleteRoutine(routineId int) error
	GetAllRoutines() ([]*Routine, error)
	FindUserById(userId int) (*User, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(user *User) (int, error)
	ConfirmUser(token string) error
	UpdateUser(user *User) (int, error)
	DeleteUser(userId int) error
	GetAllUsers() ([]*User, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

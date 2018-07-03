package models

import (
	"errors"
	"time"
)

type User struct {
	UserId    int
	Email     string
	Confirmed bool
	Hash      string
	Tag       string
	Main      string
	Bio       string
}

type Activation struct {
	ActivationId int
	UserId       int
	Token        string
	Issued       time.Time
	Expires      time.Time
	Used         time.Time
}

func (db *DB) FindUserById(userId int) (*User, error) {
	var user User

	query := "SELECT * FROM users WHERE user_id=?;"

	err := db.QueryRow(query, userId).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) FindUserByEmail(email string) (*User, error) {
	var user User

	query := "SELECT * FROM users WHERE email=?;"

	err := db.QueryRow(query, email).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) CreateUser(u *User) (int, error) {
	var userId int

	query := `INSERT INTO users(email, confirmed, hash, tag, main, bio)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING user_id;`

	err := db.QueryRow(query, u.Email, false, u.Hash, u.Tag, u.Main, u.Bio).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (db *DB) UpdateUser(u *User) (int, error) {
	var userId int

	query := `UPDATE users
		SET tag = $2, main = $3, bio = $4
		WHERE user_id = $1
		RETURNING user_id;`

	err := db.QueryRow(query, u.UserId, u.Tag, u.Main, u.Bio).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

// TODO: written after midnight, please review.
func (db *DB) ConfirmUser(userId int, token string) error {
	var a Activation

	query := `SELECT * FROM activations WHERE userId = $1;`

	err := db.QueryRow(query, userId).Scan(&a)
	if err != nil {
		return err
	}

	// check if its expired
	if token != a.Token || time.Now().After(a.Expires) {
		return errors.New("invalid token")
	}

	// switch user to confirmed
	_, err = db.Exec("UPDATE users SET confirmed = true WHERE user_id = $1;", userId)
	if err != nil {
		return err
	}

	// set token used to time.Now()
	_, err = db.Exec("UPDATE activations SET used = $1 WHERE activation_id = $2;", time.Now(), a.ActivationId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE user_id = $1;`

	_, err := db.Exec(query, userId)
	if err != nil {
		return nil
	}

	return nil
}

func (db *DB) GetAllUsers() ([]*User, error) {
	query := `SELECT * FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		u := &User{}

		err := rows.Scan(&u.UserId, &u.Email, &u.Confirmed, &u.Hash, &u.Tag, &u.Main, &u.Bio)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

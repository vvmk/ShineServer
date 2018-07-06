package models

import (
	"errors"
	"time"
)

type User struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	Confirmed bool   `json:"-"`
	Hash      string `json:"-"`
	Tag       string `json:"tag"`
	Main      string `json:"main"`
	Bio       string `json:"bio"`
}

type Activation struct {
	ActivationId int
	UserId       int
	Token        string
	Issued       time.Time
	Expires      time.Time
}

func (db *DB) FindUserById(userId int) (*User, error) {
	var u User

	query := "SELECT * FROM users WHERE user_id=$1;"

	err := db.QueryRow(query, userId).Scan(&u.UserId, &u.Email, &u.Tag, &u.Main, &u.Bio, &u.Confirmed, &u.Hash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *DB) FindUserByEmail(email string) (*User, error) {
	var u User

	query := "SELECT * FROM users WHERE email=$1;"

	err := db.QueryRow(query, email).Scan(&u.UserId, &u.Email, &u.Tag, &u.Main, &u.Bio, &u.Confirmed, &u.Hash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *DB) CreateUser(u *User) (int, error) {
	var userId int

	query := `INSERT INTO users(email, confirmed, hash, tag, main, bio)
		VALUES('$1', '$2', '$3', '$4', '$5', '$6')
		RETURNING user_id;`

	err := db.QueryRow(query, u.Email, false, u.Hash, u.Tag, u.Main, u.Bio).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (db *DB) UpdateUser(userId int, u *User) error {

	query := `UPDATE users
		SET tag = '$2', main = '$3', bio = '$4'
		WHERE user_id = $1;`

	_, err := db.Exec(query, userId, u.Tag, u.Main, u.Bio)
	if err != nil {
		return err
	}

	return nil
}

// TODO: May be necessary to add a 'reason' to the table
func (db *DB) CreateActivation(userId int, token string) error {
	now := time.Now()
	a := &Activation{
		UserId:  userId,
		Token:   token,
		Issued:  now,
		Expires: now.Add(time.Hour * 48),
	}

	query := `INSERT INTO activations(user_id, code, issued, expired)
	VALUES('$1', '$2', '$3', '$4');`

	_, err := db.Exec(query, a.UserId, a.Token, a.Issued, a.Expires)
	if err != nil {
		return err
	}

	return nil
}

// TODO: written after midnight, please review.
func (db *DB) ConfirmUser(userId int, token string) error {

	query := `SELECT * FROM activations WHERE user_id=$1;`

	var a Activation
	err := db.QueryRow(query, userId).Scan(&a.ActivationId, &a.UserId, &a.Token, &a.Issued, &a.Expires)
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

	// just delete the row. that'll invalidate'em
	_, err = db.Exec("DELETE FROM activations WHERE activation_id = $1;", a.ActivationId)
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
	query := `SELECT * FROM users;`

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

package models

type User struct {
	UserId    int
	Email     string
	Confirmed bool
	Hash      string
	Tag       string
	Main      string
	Bio       string
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
		SET tag = $3, main = $4, bio = $5
		WHERE user_id = $1
		RETURNING user_id;`

	err := db.QueryRow(query, u.UserId, u.Tag, u.Main, u.Bio).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
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

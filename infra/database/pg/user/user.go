package user

import (
	"database/sql"
	"todolist/internal/entity/user"
)

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *user.User) error {
	err := u.DB.QueryRow(`INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4);`, user.ID, user.Name, user.Email, user.Password).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u *UserDB) Delete(email string) error {
	err := u.DB.QueryRow(`DELETE FROM users WHERE email = $1;`, email).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u *UserDB) FindByEmailOrId(id, email string) (*user.User, error) {
	user := user.User{}

	row := u.DB.QueryRow(`SELECT id, name, email, password FROM users WHERE id = $1 OR email = $2;`, id, email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
